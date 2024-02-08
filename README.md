# Self-Introduction Card Generator

A hyper-minimal self-introduction card geenrator / display tool, all with a single command.
Dependencies: Standard Go libraries only.

## Usage

Simply call:

```bash
go run main
```

And this will generate a config file for you. (`config/config.json`)
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

## Edit Guide

This tool is designed to be minimal, so there aren't many complications in the tool.
You simply have to modify

- Config Construct
- Loading Attributes
- JSON generator questionnaires



#### Disclaimer
I do not know sh*t about Go, and this is not meant to be a serious, bug-free tool.
But, baby steps, right?
I would love some lessons and feedbacks in the Issues.