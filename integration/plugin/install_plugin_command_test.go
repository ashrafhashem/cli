package plugin

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("install-plugin command", func() {
	var buffer *Buffer

	BeforeEach(func() {
		helpers.RunIfExperimental("experimental until all install-plugin refactor stories are finished")
	})

	Describe("help", func() {
		Context("when the --help flag is given", func() {
			It("displays command usage to stdout", func() {
				session := helpers.CF("install-plugin", "--help")

				Eventually(session.Out).Should(Say("NAME:"))
				Eventually(session.Out).Should(Say("install-plugin - Install CLI plugin"))
				Eventually(session.Out).Should(Say("USAGE:"))
				Eventually(session.Out).Should(Say("cf install-plugin \\(LOCAL-PATH/TO/PLUGIN | URL | -r REPO_NAME PLUGIN_NAME\\) \\[-f\\]"))
				Eventually(session.Out).Should(Say("EXAMPLES:"))
				Eventually(session.Out).Should(Say("cf install-plugin ~/Downloads/plugin-foobar"))
				Eventually(session.Out).Should(Say("cf install-plugin https://example.com/plugin-foobar_linux_amd64"))
				Eventually(session.Out).Should(Say("cf install-plugin -r My-Repo plugin-echo"))
				Eventually(session.Out).Should(Say("OPTIONS:"))
				Eventually(session.Out).Should(Say("-f      Force install of plugin without confirmation"))
				Eventually(session.Out).Should(Say("-r      Name of a registered repository where the specified plugin is located"))
				Eventually(session.Out).Should(Say("SEE ALSO:"))
				Eventually(session.Out).Should(Say("add-plugin-repo, list-plugin-repos, plugins"))

				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("installing a plugin from a local file", func() {
		var pluginPath string

		BeforeEach(func() {
			pluginPath = helpers.BuildConfigurablePlugin("configurable_plugin", "some-plugin", "1.0.0",
				[]helpers.PluginCommand{
					{Name: "some-command", Help: "some-command-help"},
				},
			)
		})

		Context("when the -f flag is given", func() {
			It("installs the plugin", func() {
				session := helpers.CF("install-plugin", pluginPath, "-f")

				Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
				Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))
				Eventually(session.Out).Should(Say("Installing plugin some-plugin\\.\\.\\."))
				Eventually(session.Out).Should(Say("OK"))
				Eventually(session.Out).Should(Say("Plugin some-plugin 1\\.0\\.0 successfully installed\\."))

				Eventually(session).Should(Exit(0))

				installedPath := filepath.Join(homeDir, ".cf", "plugins", "some-plugin")

				pluginsSession := helpers.CF("plugins", "--checksum")
				expectedSha := helpers.Sha1Sum(installedPath)

				Eventually(pluginsSession.Out).Should(Say("some-plugin\\s+1.0.0\\s+%s", expectedSha))
				Eventually(pluginsSession).Should(Exit(0))

				Eventually(helpers.CF("some-command")).Should(Exit(0))

				helpSession := helpers.CF("help")
				Eventually(helpSession.Out).Should(Say("some-command"))
				Eventually(helpSession).Should(Exit(0))
			})

			Context("when the file is not executable", func() {
				BeforeEach(func() {
					Expect(os.Chmod(pluginPath, 0666)).ToNot(HaveOccurred())
				})

				It("installs the plugin", func() {
					session := helpers.CF("install-plugin", pluginPath, "-f")
					Eventually(session.Out).Should(Say("Plugin some-plugin 1\\.0\\.0 successfully installed\\."))
					Eventually(session).Should(Exit(0))

					// make sure plugin temp files are cleaned up
					pluginBinaries, err := ioutil.ReadDir(filepath.Join(homeDir, ".cf", "plugins", "temp"))
					Expect(err).ToNot(HaveOccurred())
					Expect(len(pluginBinaries)).To(Equal(0))
				})
			})

			Context("when the plugin is already installed", func() {
				BeforeEach(func() {
					Eventually(helpers.CF("install-plugin", pluginPath, "-f")).Should(Exit(0))
				})

				It("uninstalls the existing plugin and installs the plugin", func() {
					session := helpers.CF("install-plugin", pluginPath, "-f")

					Eventually(session.Out).Should(Say("Plugin some-plugin 1\\.0\\.0 is already installed\\. Uninstalling existing plugin\\.\\.\\."))
					Eventually(session.Out).Should(Say("CLI-MESSAGE-UNINSTALL"))
					Eventually(session.Out).Should(Say("Plugin some-plugin successfully uninstalled\\."))
					Eventually(session.Out).Should(Say("Plugin some-plugin 1\\.0\\.0 successfully installed\\."))

					Eventually(session).Should(Exit(0))
				})
			})

			Context("when the file does not exist", func() {
				It("tells the user that the file was not found and fails", func() {
					session := helpers.CF("install-plugin", "/some/path/that/does/not/exist", "-f")

					Eventually(session.Err).Should(Say("File not found locally, make sure the file exists at given path /some/path/that/does/not/exist"))

					Consistently(session.Out).ShouldNot(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
					Consistently(session.Out).ShouldNot(Say("Install and use plugins at your own risk\\."))

					Eventually(session).Should(Exit(1))
				})
			})

			Context("when the file is not an executable", func() {
				BeforeEach(func() {
					badPlugin, err := ioutil.TempFile("", "")
					Expect(err).ToNot(HaveOccurred())
					pluginPath = badPlugin.Name()
				})

				AfterEach(func() {
					err := os.Remove(pluginPath)
					Expect(err).ToNot(HaveOccurred())
				})

				It("tells the user that the file is not a plugin and fails", func() {
					session := helpers.CF("install-plugin", pluginPath, "-f")
					Eventually(session.Err).Should(Say("exec format error"))

					Eventually(session).Should(Exit(1))
				})
			})

			Context("when the file is not a plugin", func() {
				BeforeEach(func() {
					var err error
					pluginPath, err = Build("code.cloudfoundry.org/cli/integration/assets/non_plugin")
					Expect(err).ToNot(HaveOccurred())
				})

				It("tells the user that the file is not a plugin and fails", func() {
					session := helpers.CF("install-plugin", pluginPath, "-f")
					Eventually(session.Err).Should(Say("File %s is not a valid cf CLI plugin binary\\.", pluginPath))

					Eventually(session).Should(Exit(1))
				})
			})

			Context("command conflict", func() {
				Context("when the plugin has a command that is the same as a built-in command", func() {
					var pluginPath string

					BeforeEach(func() {
						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "some-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "version"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin some-plugin v1\\.1\\.1 could not be installed as it contains commands with names that are already used: version"))

						Eventually(session).Should(Exit(1))
					})
				})

				Context("when the plugin has a command that is the same as a built-in alias", func() {
					BeforeEach(func() {
						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "some-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "cups"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin some-plugin v1\\.1\\.1 could not be installed as it contains commands with names that are already used: cups"))

						Eventually(session).Should(Exit(1))
					})
				})

				Context("when the plugin has a command that is the same as another plugin command", func() {
					BeforeEach(func() {
						helpers.InstallConfigurablePlugin("existing-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "existing-command"},
							})

						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "new-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "existing-command"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin new-plugin v1\\.1\\.1 could not be installed as it contains commands with names that are already used: existing-command\\."))

						Eventually(session).Should(Exit(1))
					})
				})

				Context("when the plugin has a command that is the same as another plugin alias", func() {
					BeforeEach(func() {
						helpers.InstallConfigurablePlugin("existing-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "existing-command"},
							})

						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "new-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "new-command", Alias: "existing-command"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin new-plugin v1\\.1\\.1 could not be installed as it contains commands with aliases that are already used: existing-command\\."))

						Eventually(session).Should(Exit(1))
					})
				})
			})

			Context("alias conflict", func() {
				Context("when the plugin has an alias that is the same as a built-in command", func() {
					var pluginPath string

					BeforeEach(func() {
						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "some-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "some-command", Alias: "version"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin some-plugin v1\\.1\\.1 could not be installed as it contains commands with aliases that are already used: version"))

						Eventually(session).Should(Exit(1))
					})
				})

				Context("when the plugin has an alias that is the same as a built-in alias", func() {
					BeforeEach(func() {
						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "some-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "some-command", Alias: "cups"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin some-plugin v1\\.1\\.1 could not be installed as it contains commands with aliases that are already used: cups"))

						Eventually(session).Should(Exit(1))
					})
				})

				Context("when the plugin has an alias that is the same as another plugin command", func() {
					BeforeEach(func() {
						helpers.InstallConfigurablePlugin("existing-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "existing-command"},
							})

						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "new-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "new-command", Alias: "existing-command"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin new-plugin v1\\.1\\.1 could not be installed as it contains commands with aliases that are already used: existing-command\\."))

						Eventually(session).Should(Exit(1))
					})
				})

				Context("when the plugin has an alias that is the same as another plugin alias", func() {
					BeforeEach(func() {
						helpers.InstallConfigurablePlugin("existing-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "existing-command", Alias: "existing-alias"},
							})

						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "new-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "new-command", Alias: "existing-alias"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin new-plugin v1\\.1\\.1 could not be installed as it contains commands with aliases that are already used: existing-alias\\."))

						Eventually(session).Should(Exit(1))
					})
				})
			})

			Context("alias and command conflicts", func() {
				Context("when the plugin has a command and an alias that are both taken by another plugin", func() {
					BeforeEach(func() {
						helpers.InstallConfigurablePlugin("existing-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "existing-command", Alias: "existing-alias"},
							})

						pluginPath = helpers.BuildConfigurablePlugin(
							"configurable_plugin", "new-plugin", "1.1.1",
							[]helpers.PluginCommand{
								{Name: "existing-command", Alias: "existing-alias"},
							})
					})

					It("tells the user about the conflict and fails", func() {
						session := helpers.CF("install-plugin", "-f", pluginPath)

						Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
						Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin new-plugin v1\\.1\\.1 could not be installed as it contains commands with names and aliases that are already used: existing-command, existing-alias\\."))

						Eventually(session).Should(Exit(1))
					})
				})
			})
		})

		Context("when the -f flag is not given", func() {
			Context("when the user says yes", func() {
				BeforeEach(func() {
					buffer = NewBuffer()
					buffer.Write([]byte("y\n"))
				})

				It("installs the plugin", func() {
					session := helpers.CFWithStdin(buffer, "install-plugin", pluginPath)

					Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
					Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))
					Eventually(session.Out).Should(Say("Do you want to install the plugin %s\\? \\[yN\\]: y", pluginPath))
					Eventually(session.Out).Should(Say("Installing plugin some-plugin\\.\\.\\."))
					Eventually(session.Out).Should(Say("OK"))
					Eventually(session.Out).Should(Say("Plugin some-plugin 1\\.0\\.0 successfully installed\\."))

					Eventually(session).Should(Exit(0))

					pluginsSession := helpers.CF("plugins", "--checksum")
					expectedSha := helpers.Sha1Sum(
						filepath.Join(homeDir, ".cf/plugins/some-plugin"))
					Eventually(pluginsSession.Out).Should(Say("some-plugin\\s+1.0.0\\s+%s", expectedSha))
					Eventually(pluginsSession).Should(Exit(0))

					Eventually(helpers.CF("some-command")).Should(Exit(0))

					helpSession := helpers.CF("help")
					Eventually(helpSession.Out).Should(Say("some-command"))
					Eventually(helpSession).Should(Exit(0))
				})

				Context("when the plugin is already installed", func() {
					BeforeEach(func() {
						Eventually(helpers.CF("install-plugin", pluginPath, "-f")).Should(Exit(0))
					})

					It("fails and tells the user how to force a reinstall", func() {
						session := helpers.CFWithStdin(buffer, "install-plugin", pluginPath)

						Eventually(session.Out).Should(Say("FAILED"))
						Eventually(session.Err).Should(Say("Plugin some-plugin 1\\.0\\.0 could not be installed\\. A plugin with that name is already installed\\."))
						Eventually(session.Err).Should(Say("TIP: Use 'cf install-plugin %s -f' to force a reinstall\\.", pluginPath))

						Eventually(session).Should(Exit(1))
					})
				})
			})

			Context("when the user says no", func() {
				BeforeEach(func() {
					buffer = NewBuffer()
					buffer.Write([]byte("n\n"))
				})

				It("does not install the plugin", func() {
					session := helpers.CFWithStdin(buffer, "install-plugin", pluginPath)

					Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
					Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))
					Eventually(session.Out).Should(Say("Do you want to install the plugin %s\\? \\[yN\\]: n", pluginPath))
					Eventually(session.Err).Should(Say("Plugin installation cancelled"))
					Eventually(session.Out).Should(Say("FAILED"))

					Eventually(session).Should(Exit(1))
				})

				Context("when the plugin is already installed", func() {
					BeforeEach(func() {
						Eventually(helpers.CF("install-plugin", pluginPath, "-f")).Should(Exit(0))
					})

					It("does not uninstall the existing plugin", func() {
						session := helpers.CFWithStdin(buffer, "install-plugin", pluginPath)

						Eventually(session.Err).Should(Say("Plugin installation cancelled"))

						Consistently(session.Out).ShouldNot(Say("Plugin some-plugin 1\\.0\\.0 is already installed\\. Uninstalling existing plugin\\.\\.\\."))
						Consistently(session.Out).ShouldNot(Say("CLI-MESSAGE-UNINSTALL"))
						Consistently(session.Out).ShouldNot(Say("Plugin some-plugin successfully uninstalled\\."))

						Eventually(session).Should(Exit(1))
					})
				})
			})

			Context("when the user interrupts with control-c", func() {
				BeforeEach(func() {
					buffer = NewBuffer()
					buffer.Write([]byte("y")) // but not enter
				})

				It("does not install the plugin and does not create a bad state", func() {
					session := helpers.CFWithStdin(buffer, "install-plugin", pluginPath)

					Eventually(session.Out).Should(Say("Attention: Plugins are binaries written by potentially untrusted authors\\."))
					Eventually(session.Out).Should(Say("Install and use plugins at your own risk\\."))
					Eventually(session.Out).Should(Say("Do you want to install the plugin %s\\? \\[yN\\]:", pluginPath))

					session.Interrupt()

					Eventually(session.Out).Should(Say("FAILED"))

					Eventually(session).Should(Exit(1))

					// make sure cf plugins did not break
					Eventually(helpers.CF("plugins", "--checksum")).Should(Exit(0))

					// make sure a retry of the plugin install works
					retrySession := helpers.CF("install-plugin", pluginPath, "-f")
					Eventually(retrySession.Out).Should(Say("Plugin some-plugin 1\\.0\\.0 successfully installed\\."))
					Eventually(retrySession).Should(Exit(0))
				})
			})
		})
	})
})
