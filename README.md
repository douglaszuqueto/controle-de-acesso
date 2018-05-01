# Controle de Acesso

**Toda documentação do projeto encontra-se em fase de desenvolvimento!**

![img](https://raw.githubusercontent.com/douglaszuqueto/controle-de-acesso/master/.github/images/print-dashboard.png)

## Índice

- [Introdução](#introdução)
- [Objetivo](#objetivo)
- [Materiais](#materiais)
- [Tecnologias](#tecnologias)
- [Arquitetura](#arquitetura)
- [Organização](#organização)
- [Fluxo](#fluxo)
- [Rodando o projeto](#rodando-o-projeto)
- [Resultados](#resultados)
- [Referências](#referências)


## Introdução

Depois de completar quase 1 ano que lancei o primeiro 'projeto' de envolvendo controle de acesso, eis que foi desenvolvido uma versão mais completa.

Para quem ainda não conhece, o projeto [Controle de acesso utilizando NodeMCU, RFID, MQTT e Banco de Dados MySQL](https://github.com/douglaszuqueto/esp8266-rfid-banco-de-dados) foi desenvolvido em Julho de 2017, onde tinha como objetivo desenvolver um projeto simples e direto ao ponto para sanar muitas dúvidas que membros da comunidade tinham, principalmente no que tange integração, arquitetura e comunicação com banco de dados. Portanto, para quem ainda não viu, eu recomendo dar uma olhada no projeto antigo para já ter um embasamento do que se trata - [link](https://github.com/douglaszuqueto/esp8266-rfid-banco-de-dados).

Como o objetivo era desenvolver um projeto mais completo, e como também, eu sempre estou em busca de testar e validar novas tecnologias, neste projeto não foi diferente - Utilizei **Go(golang)** como linguagem de programação para o *Back-end* e **VueJS** como biblioteca para construção do *Front-end*. No embarcado fiquei com as mesmas tecnologias pois é a que mais domino no momento.

## Objetivo

Desenvolver tal projeto que seja de fácil implementação - tanto no que tange **Embarcado** como também **Arquitetura** e **Software** e também oferecer uma boa experiência através de networking e troca de ideias - sendo os 2 últimos itens mencionados, os fatores que mais levo em consideração no desenvolvimento de algum projeto :)

## Materiais

Abaixo segue uma seção resumida dos materiais principais que foram utilizadas no decorrer do desenvolvimento do projeto.

* Embarcado
    * NodeMCU
    * Rfid Mfrc522 Mifare
* Servidor
    * Raspberry Pi 3
    * Raspberry Pi Zero W
    * Servidor em Nuvem

**Observação:** No que tange servidor, apenas um item citado acima basta! Detalharei com mais calma em outro tópico.

## Tecnologias

* Firmware
    * [PubSubClient](https://github.com/knolleary/pubsubclient)
    * [MFRC522](https://github.com/miguelbalboa/rfid/)
* Back-end - Go
    * [pq - A pure Go postgres driver for Go's database/sql package](https://github.com/lib/pq)
    * [gorilla/handlers](https://github.com/gorilla/handlers)
    * [gorilla/mux](https://github.com/gorilla/mux)
    * [Eclipse Paho MQTT Go client](https://github.com/eclipse/paho.mqtt.golang)
    * [jwt-go](https://github.com/dgrijalva/jwt-go)
    * [statik](https://github.com/rakyll/statik)
* Front-end - VueJS
    * [VueJS](https://github.com/vuejs/vue)
    * [Vue Router](https://github.com/vuejs/vue-router)
    * [Bootstrap 4](https://github.com/twbs/bootstrap)
    * [Axios](https://github.com/axios/axios)
    * [MQTT.js](https://github.com/mqttjs/MQTT.js)
    * [Font-Awesome](https://github.com/FortAwesome/Font-Awesome)
    * [Sweet Alert 2](https://github.com/sweetalert2/sweetalert2)

## Arquitetura

inserir diagrama geral aqui

## Organização

**Level 1**
```
~/projetos/controle-de-acesso on  master ⌚ 14:43:36
$ tree -L 1
.
├── app
├── bin
├── cli
├── database
├── firmware
├── front
├── README.md
└── systemd

7 directories, 1 file
```

Num primeiro momento, o projeto é formado por **7 pastas** - podendo crescer futuramente. 

**Observação:** Detalhes mais 'aprofundados' de cada pasta será abordado em um **README** dentro do sua respectiva pasta.

### 1 - app

Conterá todo back-end da aplicação, ou seja, tudo que é necessário para realizar o build da aplicação e 'colocar em produção'. Possui 3 scripts(em bash mesmo) para colocar todo o projeto rodar em 3 ambientes: **Docker**, **Local** e na **Raspberry Pi**.

### 2 - bin

Pasta responsável por armazenar os binários(aplicação compilada) para arquitetura **Linux 64** e **ARM(raspberry)**

### 3 - cli

Binário que será desenvolvido para manipular todos *endpoints* da aplicação através da linha de comando.

**Ex:**

* 1º Criar uma tag
* 2º Associar a um usuário
* 3º Associar a um device

**Observação: Em desenvolvimento**

### 4 - database

O banco de dados utilizado foi o **PostgreSQL** - um banco que fazia algum tempo que queria validar em 1 projeto :)

Confira abaixo a estrutura do projeto CDA.

![img](https://raw.githubusercontent.com/douglaszuqueto/controle-de-acesso/master/database/estrutura.png)

Em resumo, o projeto se dá pela seguinte forma:

* administradores para acessar o painel
* usuários que por sua vez podem possuir diversas tags
* devices - podem ser catracas, portas de acesso dentre outros cenários
* tais tags são associadas a devices, assim liberando ou proibindo o acesso daquela tag específica e que pertence a um usuario X

Demais informações serão descritas na respectiva pasta *database*

### 5 - firmware

A pasta firmware contem 2 subpastas, são elas: **cda** e **reader**.

O firmware **cda** é o principal. Ele que rodará no seu hardware principal - catracas, portas de acesso e etc.

Já o firmware **reader** é como se fosse um *Plus*, você poderá ter um hardware adicional no balcão - por exemplo, para realizar o cadastro de tags. Então este firmware em conjunto com um nodemcu e um leitor rfid irá preencher automaticamente o **uuid** quando a tag for lida na tela de cadastro da Dashboard.

### 6 - front

### 7 - systemd

O **systemd** será utilizado quando seu sistema operacional tiver suporte e você queira que, ao ligar/reiniciar seu servidor o back-end seja iniciado automaticamente, ou, em caso de falhas, o mesmo seja reiniciado.

**Lista de diretórios mais detalhada**
```
~/projetos/controle-de-acesso on  master! ⌚ 14:44:44
$ tree -L 2
.
├── app
│   ├── app
│   ├── bin
│   ├── config
│   ├── deploy-docker.sh
│   ├── deploy-local.sh
│   ├── deploy-rpi.sh
│   ├── Dockerfile
│   ├── handlers
│   ├── main.go
│   ├── models
│   ├── mqtt
│   ├── pkg
│   ├── public
│   ├── README.md
│   ├── routes.go
│   ├── services
│   ├── src
│   ├── statik
│   └── utils
├── bin
│   ├── app
│   └── README.md
├── cli
│   └── README.md
├── database
│   ├── controle-de-acesso.mwb
│   ├── controle-de-acesso.mwb.bak
│   ├── estrutura.png
│   ├── README.md
│   └── sistema.sql
├── firmware
│   ├── cda
│   └── reader
├── front
│   ├── build
│   ├── config
│   ├── dist
│   ├── Dockerfile
│   ├── index.html
│   ├── node_modules
│   ├── package.json
│   ├── README.md
│   ├── src
│   ├── static
│   └── yarn.lock
├── README.md
└── systemd
    ├── cda
    ├── cda.service
    └── README.md

26 directories, 25 files
```

## Fluxo

## Rodando o projeto

Neste tópico, tudo que abordar Sistema Operacional, estarei me baseando em **Debian/Ubuntu**, mesmo tendo como distro principal o Fedora. Tudo isso pois como o foco de servidor é na Raspberry Pi, sabemos que ela é baseada no Debian, então facilita tudo :).

Mas se você quer rodar numa distro diferente, não tem problemas, o que mais muda mesmo é o **gerenciador de pacotes**, de resto é tranquilo.

Saliento também, que terá 2 ambientes, ou seja, tem um ambiente no qual o build é gerado por você e o cenário no qual você pode aproveitar os binários que estarão na **pasta bin**.

### Preparando o ambiente

Para começar, tenha seu ambiente atualizado :)

```
sudo apt-get update
sudo apt-get upgrade

sudo apt-get install vim git wget
```

### Broker MQTT - Mosquitto

```
sudo apt-get install mosquitto
```

No diretório **/etc/mosquitto/conf.d** crie o arquivo *cda.conf*

```
sudo vim /etc/mosquitto/conf.d/cda.conf
```

Adicione o seguinte conteudo dentro do arquivo em um editor de sua preferência:

```
# Configurations

allow_anonymous true

log_type error
log_type warning
log_type notice
log_type information
log_type debug

# MQTT

listener 1883

# MQTT Websockets

listener 8083
protocol websockets
```

Caso você queira utilizar portas diferentes fique a vontade para mudar. Porém certifique-se que elas não estarão sendo utilizadas por outro serviço.

Para atualizar as configurações do mosquitto você precisa dar um restart no daemon do systemd e em seguida reiniciar o serviço do mosquitto.

```
sudo systemd daemon-reload
sudo systemd restart mosquitto
```

### Banco de dados - PostgreSQL

Para realizar a instalação, certifique-se que esteja tudo certo com o pacote. Para isso realize uma busca pelo mesmo.

```
sudo apt-cache search postgresql-9.6
```

Sua saída deverá ter um pacote exatamente com o nome pesquisado acima, e também aparecerá outros plugins. Para instalar o banco de dados, basta executar o comando abaixo:

```
sudo apt-get install postgresql-9.6
```

#### Configurações adicionais

##### Liberando acesso para a rede
No arquivo **/etc/postgresql/9.6/main/postgresql.conf**, procure pela linha que contenha esta nomenclatura: **listen_addresses**. Independente do conteúdo, troque pela seguinte configuração:

```
listen_addresses = '*'
```

##### Setando timezone
No mesmo arquivo anterior, busque por **timezone**, e troque seu conteúdo por:

```
timezone = 'America/Sao_Paulo'
```

**Obeservação:** Caso seja de uma timezone diferente, não esqueça de mudar.

##### Adicionando regra de autenticação criptografada

No final do arquivo **/etc/postgresql/9.6/main/pg_hba.conf** adicione a seguinte linha:

```
host	all		all		192.168.0.1/24		md5
```

**Observação:** Não esqueça de mudar sua rede caso a mesma esteja em uma faixa diferente. No meu caso, está na rede 0: 192.168.**0**.1/24

##### Setando senha do usuário default - postgres

```
# 1º passo
sudo -u postgres psql

# 2º passo
ALTER USER postgres PASSWORD = 'root'

# 3º passo
\q

# 4º passo
sudo systemctl restart postgresql

```

### Realizando build do projeto

### Colocando rodar pela 1º vez

### Configurando para iniciar com o Sistema Operacional

## Resultados

## Referências

- [Controle de acesso utilizando NodeMCU, RFID, MQTT e Banco de Dados MySQL](https://github.com/douglaszuqueto/esp8266-rfid-banco-de-dados)

# Gostou do projeto?

Se você curtiu o projeto e quer trocar uma ideia, foi criado um grupo no **Telegram** para debate! - [clique aqui para participar](https://t.me/joinchat/BPOe2hAc3mNcw_y7f0qipg).