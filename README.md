# Software Engineer Test solution

- Author: `Kaique Garcia Menezes`
- Contact: [`contato@kaiquegarcia.dev`](mailto:contato@kaiquegarcia.dev)
- Check my profile: [kaiquegarcia.dev](https://kaiquegarcia.dev)
- LinkedIn: [@kaiquegarcia](https://www.linkedin.com/in/kaiquegarcia/)

# Prerequisites

- [Golang](https://go.dev/dl/) v1.17;

# How to use

1. Clone the repository
2. Run `go run . report` to call the program.

# Customizations

You can send some arguments to customize the program and execute more cool stuff.

1. Group by mean: just include the argument `--group-by mean`;
2. Consider `<world>` as a player: `--world-as-player true`;
3. Change the log's file path to run the program with other logs: `--filepath <absolutePath>`.

Example:
```
go run . report --group-by mean --world-as-player false --filepath qgames.log
```

You can also see the `help` message with `go run . help`.