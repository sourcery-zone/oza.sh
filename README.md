# ⚡ `oza.sh`: Fast, Lazy, Realtime Shell Context Engine

**oza.sh** is a fast, and extensible command-line tool for building
feature-rich status lines in terminal multiplexers (like `tmux`);
and maybe shells too!

It's like a hyper-efficient, modular `powerline` or `starship`,
without being bound to the limitations of your shell!

> 🌿 Git status? 🔶 AWS profile? ⛅ Kube context? All fetched *only
  when needed*, in parallel, and optionally cached.

---

## 🚀 Why oza.sh?

- ✅ **Lazy evaluation**: snappy fast and real-time.
- 🚀 **Parallel collectors**: Git, AWS, Kube, and more run concurrently.
- 💡 **Shell-friendly**: safe for prompts, placeholders, and caching.
- 🧠 **Real-time in tmux**: perfect for always-on status bars.
- 🛠️ **Extensible in Go**: build your own collectors with minimal effort.
- 🔐 **Remote-aware**: can read pane environments over SSH.
- 🗂️ **Context-aware**: only display what you need in your current context.

---

## 📦 Example

A tmux-friendly command to show Git + AWS status only if needed:

```bash
oza --format '{{ .Git.Branch }}{{ if .Git.Dirty }}*{{ end }}{{ .AWS.Profile }}'
```
## 🧰 Supported Collectors

| Collector | Example Output                                 | Notes                          |
| --------: | ---------------------------------------------- | ------------------------------ |
|       Git | `main*`                                        | Branch + dirty status          |
|       AWS | `dev-admin`                                    | Uses `AWS_PROFILE`, lazy env   |
|      Kube | `staging/default`                              | Optional, for K8s-aware shells |
|    Future | 🤖 Python, Terraform, custom ones (via plugin) |                                |

## Installation

[TBD]

## 🧠 How It Works

- You provide a template like `{{ .Git.Branch }}` by one of the
  following manners:
  - Your global session (`tmux` pane, shell prompt, etc.)
  - Current folder
  - Current Git repository
- `oza` analyzes the format, sees `.Git` is needed
- The daemon will start fetching the relevant data in relevant interfals
- `oza` won't however wait, and until the data is present, show a placeholder
- When the data become available, it'll be displayed where you need it

## 📘 Name Meaning

ozash (اوضاعش in Persian for “its status” or “its situation”) Neatly
describes what this tool does without blocking or guessing. 😉

## 📅 Roadmap

- [x] Git collector
- [ ] AWS collector
- [x] Template format engine
- [ ] Background daemon mode
- [ ] Shell-safe cache fallback
- [ ] SSH-aware environment sync
- [ ] Plugin system for custom collectors

## 🤝 Contributing

Pull requests welcome! Want to add a new collector? Check out
internal/collectors/ for examples.

