# Billing Dashboard API

Dibuat menggunakan bahasa pemrograman Go. Digunakan untuk interkoneksi dengan dashboard billing system.

## Development

```bash
compiledaemon --command="./go-billing-dashboard-api"
```

## Deployment

```bash
docker build -t go-billing-dashboard-api .
docker container create --name go-billing-dashboard-api_1 -p 7720:7720 go-billing-dashboard-api
docker start go-billing-dashboard-api_1
```

Rebuild

```bash
sudo docker stop go-billing-dashboard-api_1 && sudo docker rm go-billing-dashboard-api_1 && sudo docker rmi go-billing-dashboard-api && sudo docker build -t go-billing-dashboard-api . && sudo docker container create --name go-billing-dashboard-api_1 -p 7720:7720 go-billing-dashboard-api && sudo docker start go-billing-dashboard-api_1
```

## Lisensi

Hak Cipta dilindungi dan milik @atozpw
