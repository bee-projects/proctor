# Overview

While forwarwarding traffic between systems, Proctor can facilitate the role of an authentication broker. 

![authentication](./images/authentication.png)

An  *Authenticator plugin* is an independant natively compiled binary that exposes an API over net/rpc. This allows us to generecise any given authentication mechanism using a plug-in model and enables its reuse across different applications and environments.


