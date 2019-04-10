# d2prox

Diablo II proxy written in Go. Proxying allows the user to bypass IP rate limits set by battle.net. 

*It's currenly hardcoded for the European realm. Additionally it will only run on the same machine as the Diablo client, which makes it completely useless. However, anyone who knows how to compile it should be able to configure it to work with any realm/ip.*

Since the code allows inspection and modification of any D2 packet, it could theoretically be used as a base for packet editing, packet filtering and implementing packet based hacks & bots.

## Disclaimer

This software is unfinished and not properly tested. It's only meant as a proof of concept. However, it does work and could be used with minimal modification. As far as I know, it should be undetectable. 

## How to use

* Rewrite ``europe.battle.net`` to ``127.0.0.1`` in your systems hosts file
* Run d2prox
* Login and play

## Method

It works by implementing 3 separate proxies for Battle.net, the Diablo 2 Realm Server and the Diablo 2 Game Server. Application-level proxying is achieved by deep packet inspection and modification.

The basic idea is as follows:

* Configure the system to resolve ``_realm_.battle.net`` to the proxy ip
* Client connects to our battle.net proxy
* Once the player authenticates, battle.net sends the realm server ip to the client (``SID_REALMLOGONEX``). This packet is intercepted by the proxy, and replaced with ``127.0.0.1``. The original realm server ip is stored in a map which correlates clients (using the MCP tokens) to realm server ips.
* Client connects to the realm server proxy
* Client sends ``MCP_STARTUP`` containing its MCP token. The MCP token is used to retrieve the original realm server ip, and a connection is opened to it. From now on, all traffic is forwarded to the real realm server.
* Client creates or joins a game. ``MCP_JOINGAME`` is intercepted, and the game server ip is also replaced with ``127.0.0.1``, and a token is used to map the client to this original ip.
* Client connects to the game server proxy
* Client sends ``D2GS_GAMELOGON`` with the game token, which is used to retrieve the original ip and connect to the game server. From now on, all traffic is forwarded to the real game server.