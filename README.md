# Beibootprojekt Installation per Docker

### Repository clonen und in Verzeichnis wechseln

```
git clone https://github.com/mi-classroom/mi-web-technologien-beiboot-ss2020-p2sk
cd PathTo/mi-web-technologien-beiboot-ss2020-p2sk
```

### Image bauen

```
docker-compose build
```

### Container starten

```
docker-compose up -d
```

### Im Browser aufrufen

```
localhost:8080
```

### Container beenden

```
docker-compose down [--remove-orphans]
```