# Mimo Counter

Mimo Counter is a simple adapter for MimoLive and Bitfocus Companion for Elgato
Stream Deck. It allows the Companion to treat a Text Layer in Mimo Live as a
manual counter. Companion talks to Mimo Counter via simple TCP commands, and
Mimo Counter updates the Mimo Live layer via its HTTP API.

Built for MimoLive 6.3.1 and Companion 3.0.0.

## Configuration

```bash
export MIMO_HOST=localhost
export MIMO_PORT=8989
export MIMO_DOCUMENT_ID=327972873
export MIMO_LAYER_ID=2FA3DFAA-5931-41DF-9B96-FB5518445A7F

export TCP_HOST=localhost
export TCP_PORT=6000
```

## TCP API

- `get`
- `set $n`
- `inc`
- `dec`
