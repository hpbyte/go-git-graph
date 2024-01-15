# go-git-graph

`go-git-graph` is a tool for visualizing contribution stats from your local git repositories.

We often have separate **GitHub** accounts for personal and work, or have repositories across various platforms like **BitBucket** or **GitLab**. This can make it tricky to get a unified view of all our contributions. But in the end, our local machine is a **single-source-of-truth** for all of our contributions, so that's why this tool can be used to visualize in a unified way.

![Screenshot](./assets/screenshot.png?raw=true "GoGitGraph")

## Installation

To install `go-git-graph`, follow these steps:

```bash
git clone https://github.com/hpbyte/go-git-graph

cd go-git-graph/cmd/go-git-graph

go build -o ggg
```

## Configuration

Please create a config file named `.ggg.conf.toml` in your home directory as follows:

```bash
cd ~

# pls update with your git users' emails
echo -e "[user]\nemails = [\"user1@example.com\", \"user2@example.com\", \"user3@example.com\"]" > .ggg.conf.toml
```

### Usage

After building the project, you can run it using:

```bash
./ggg
```

Help is also available via:

```bash
./ggg -h
```

By default, it scans in the current directory. In order to scan in a specific directory:

```bash
./ggg -p /path/to/dir
```

By default, stats are calculated for the current year. To see the stats by a specific year:

```bash
./gggg -p /path/to/dir -y 2023
```

Scanned directories are always cached, you can clear this cache by:

```bash
./ggg -p /path/to/dir -c true
```
