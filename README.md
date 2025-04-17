---

# 🔍 buscarIps

Aplicação de linha de comando feita em Go que permite buscar endereços IP de hosts na internet usando a biblioteca [`urfave/cli`](https://github.com/urfave/cli).

## 🚀 Funcionalidades

- Buscar IPs de um domínio/host usando DNS (`net.LookupIP`)

## 🛠️ Tecnologias

- [Go](https://golang.org/)
- [urfave/cli](https://github.com/urfave/cli) - Framework para CLI apps em Go

## 📦 Instalação

Clone o repositório:

```bash
git clone https://github.com/svitorz/buscarIps.git
cd buscarIps
```

## ▶️ Uso

Para executar a aplicação, utilize o comando:

```bash
go run main.go ip --host amazon.com.br
```

Se nenhum `--host` for passado, será utilizado o valor padrão `devbook.com.br`.

### Exemplo de saída:

```
54.239.28.85
54.239.28.102
```

## 📁 Estrutura

- `main.go`: ponto de entrada da aplicação
- `app/`: contém a definição dos comandos CLI e ações

## 📄 Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---
