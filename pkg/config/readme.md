# Config Package

Locates and parses a chia configuration file into a config struct. If the `CHIA_ROOT` environment variable is set, the config will be loaded from that location. Otherwise, the package will look in `~/.chia/mainnet`. [See the wiki for for more information on using the `CHIA_ROOT` variable.](https://github.com/Chia-Network/chia-blockchain/wiki/INSTALL#testnets)
