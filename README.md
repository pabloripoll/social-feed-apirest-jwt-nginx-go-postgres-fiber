<div id="top-header" style="with:100%;height:auto;text-align:right;">
    <img src="./static/files/pr-banner-long.png">
</div>

# SOCIAL FEED - GO - Fiber v2 - Gorm / Postgre 16.4

This repository contains a basic example of a RESTful API service built with **GO - Fiber V2**, intended for research purposes and as a demonstration of my developer profile. It implements the core features of a minimal, custom social feed application and serves as a reference project for learning, experimentation, or as a back-end development code sample.

> ⚠️ **Project Status: In Development**
>
> This repository is currently under active development and not yet ready for use. Features and APIs may change, and breaking changes can occur. Please, be aware that the codebase is actively evolving.

## Project Overview

The API supports a registry of platform "members," enabling users to create posts and voting with like or dislike other users' posts. An administrator role is provided for managing users, content, and platform statistics via a dedicated back office.

## Content of this page:

- [Infrastructure Platform](#infrastructure-platform)
- [REST API - GO Fiber](#apirest-setup)
<br><br>

## <a id="infrastructure-platform"></a>Infrastructure Platform

You can use your own local infrastructure to clone and run this repository. However, if you use [GNU Make](https://www.gnu.org/software/make/) installed, we recommend using the dedicated Docker repository [**NGINX 1.28, GO 1.25 - POSTGRES 16.4**](https://github.com/pabloripoll/docker-platform-nginx-go-1.25-pgsql-16.4)

With just a few configuration steps, you can quickly set up this project—or any other—with this same required stack.

**Repository directories structure overview:**
```
.
├── apirest (GO - Fiber)
│   ├── app
│   ├── bootstrap
│   ├── main.go
│   └── ...
│
├── platform
│   ├── nginx-go
│   │   ├── docker
│   │   │   ├── config
│   │   │   │   ├── go
│   │   │   │   ├── nginx
│   │   │   │   └── supervisor
│   │   │   ├── .env
│   │   │   ├── docker-compose.yml
│   │   │   └── Dockerfile
│   │   │
│   │   └── Makefile
│   └── postgres-16.4
│       ├── docker
│       └── Makefile
├── .env
├── Makefile
└── README.md
```

Follow the documentation to implement it:
- https://github.com/pabloripoll/docker-platform-nginx-go-1.25-pgsql-16.4?tab=readme-ov-file#api-service-container-settings
<br><br>

## <a id="apirest-setup"></a>REST API - GO Fiber

The following steps assume you are using the recommended [NGINX-GO with Postgres 16.4 platform repository](https://github.com/pabloripoll/docker-platform-nginx-go-1.25-pgsql-16.4).

Remove the platform repository `./apirest` content if any. The `.gitkeep` added after cloning the API repository.

Placed in platform repository root, remove `./apirest` from git tracking
```bash
$ git rm -r --cached -- "apirest/*" ":(exclude)apirest/.gitkeep"
```

Clone the API repository
```bash
$ cd ./apirest
$ git clone https://github.com/pabloripoll/social-feed-apirest-jwt-nginx-go-postgres-fiber .
$ touch .gitkeep
```
<br>

Set up environment
- Copy `.env.example` to `.env` and adjust settings (database, JWT secret, etc.)
<br>

Access container to install the project. Once accessed into the container, you will be placed into root proyect directory at `/var/www`
```bash
# outside container
$ make apirest-ssh

# inside container
/var/www $
```

Install project dependencies
```bash
/var/www $ go mod tidy
```
<br>

Ensure that connection to database is correct through any database client or by watching database platform docker logs.

Build and run the application
```bash
/var/www $ go run main.go
```

Once the application is running, it is not neccessary to keep tty connected. Just cancel and restart the Supervisord GO service that is as to be set at `./platform/nginx-go/docker/config/supervisor/conf.d/go.conf` *(remember that this service name can be customizable)*
```bash
/var/www $ sudo supervisorctl restart go
```
<br><br>

## Contributing

Contributions are very welcome! Please open issues or submit PRs for improvements, new features, or bug fixes.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -am 'feat: Add new feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Create a new Pull Request
<br><br>

## License

This project is open-sourced under the [MIT license](LICENSE).

<!-- FOOTER -->
<br>

---

<br>

- [GO TOP ⮙](#top-header)

<div style="with:100%;height:auto;text-align:right;">
    <img src="./static/files/pr-banner-long.png">
</div>