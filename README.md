# Self-Introduction Card Generator

A hyper-minimal self-introduction card geenrator / display tool, all with a single command.
Dependencies: Standard Go libraries only.

## Usage

Simply call:

```bash
go run main
```

And this will generate a config file for you.
You can of course create or modify the config file manually.

Now, run:

```bash
go run main
```

again, and this will display the self-introduction card for you.

If you need to modify the config file, you can do so by:

- Modifying the `config.json` file directly
- Or, you can use the following command to modify the config file:

```bash
go run main --edit
```
