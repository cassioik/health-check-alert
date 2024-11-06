# Health Check Alert

Este projeto é um serviço de verificação de saúde que envia alertas para um webhook do Discord quando a verificação falha.

## Como rodar o projeto usando a versão `latest` disponível no Docker Hub

1. **Certifique-se de que o Docker está instalado**: Verifique se o Docker está instalado na sua máquina rodando `docker --version`. Se não estiver instalado, siga as instruções em [docker.com](https://docs.docker.com/get-docker/).

2. **Execute a imagem Docker**: Use o comando `docker run` para executar a imagem, passando as variáveis de ambiente necessárias (`URL_TO_PING`, `DISCORD_WEBHOOK_URL` e `PING_INTERVAL`):

    ```sh
    docker run -e URL_TO_PING="http://example.com" -e DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/..." -e PING_INTERVAL=5 docker.io/cassioik/health-check-alert:latest
    ```

## Como fazer alterações e subir uma nova versão para o Docker Hub

1. **Clone o repositório**:

    ```sh
    git clone https://github.com/cassioik/health-check-alert.git
    cd health-check-alert
    ```

2. **Faça as alterações necessárias no código**.

3. **Atualize a versão no GitHub**: Crie uma nova tag de versão e faça o push para o GitHub:

    ```sh
    git tag v1.0.1
    git push origin v1.0.1
    ```

4. **GitHub Actions**: O GitHub Actions está configurado para construir e fazer o push da imagem Docker automaticamente quando uma nova tag é criada. O workflow `docker-publish.yml` cuidará do build e push da imagem para o Docker Hub.

## Estrutura do Projeto

- `main.go`: Código principal do serviço de verificação de saúde.
- `Dockerfile`: Arquivo Docker para construir a imagem do serviço.
- `go.mod` e `go.sum`: Arquivos de dependências do Go.
- `.github/workflows/docker-publish.yml`: Workflow do GitHub Actions para construir e fazer o push da imagem Docker.

## Variáveis de Ambiente

- `URL_TO_PING`: URL que será verificada.
- `DISCORD_WEBHOOK_URL`: URL do webhook do Discord para enviar alertas.
- `PING_INTERVAL`: Intervalo de tempo (em minutos) entre as verificações.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.