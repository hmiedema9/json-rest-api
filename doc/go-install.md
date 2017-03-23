

## Go Installation
If you are upgrading from an older version of Go you must first remove the existing version.

### Linux, Mac OS X, and FreeBSD tarballs
Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:

`tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`


Choose the archive file appropriate for your installation. For instance, if you are installing Go version 1.2.1 for 64-bit x86 on Linux, the archive you want is called go1.2.1.linux-amd64.tar.gz.

_(Typically these commands must be run as root or through sudo.)_

Add `/usr/local/go/bin` to the PATH environment variable. You can do this by adding this line to your `/etc/profile` (for a system-wide installation) or `$HOME/.profile`:

`export PATH=$PATH:/usr/local/go/bin`

### Installing to a custom location

The Go binary distributions assume they will be installed in /usr/local/go (or c:\Go under Windows), but it is possible to install the Go tools to a different location. In this case you must set the GOROOT environment variable to point to the directory in which it was installed.

For example, if you installed Go to your home directory you should add commands like the following to `$HOME/.profile`:

`export GOROOT=$HOME/go1.X`

`export PATH=$PATH:$GOROOT/bin`

_Note: GOROOT must be set only when installing to a custom location._

### Mac OS X package installer
Download the package file, open it, and follow the prompts to install the Go tools. The package installs the Go distribution to `/usr/local/go`.

The package should put the `/usr/local/go/bin` directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.

### Windows
The Go project provides two installation options for Windows users (besides installing from source): a zip archive that requires you to set some environment variables and an MSI installer that configures your installation automatically.

### MSI installer

Open the MSI file and follow the prompts to install the Go tools. By default, the installer puts the Go distribution in `c:\Go`.

The installer should put the `c:\Go\bin` directory in your PATH environment variable. You may need to restart any open command prompts for the change to take effect.

### Zip archive

Download the zip file and extract it into the directory of your choice (we suggest `c:\Go`).

If you chose a directory other than `c:\Go`, you must set the GOROOT environment variable to your chosen path.

Add the bin subdirectory of your Go root (for example, `c:\Go\bin`) to your PATH environment variable.

### Setting environment variables under Windows

Under Windows, you may set environment variables through the "Environment Variables" button on the "Advanced" tab of the "System" control panel. Some versions of Windows provide this control panel through the "Advanced System Settings" option inside the "System" control panel.
