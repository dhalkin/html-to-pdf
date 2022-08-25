# html-to-pdf

Is a simple web application for converting html to pdf

### Install

1. Clone repository to local computer:

```bash
git clone git@github.com:dhalkin/html-to-pdf.git
```

2. Build application image

``` bash
docker-compose build
```

3. Change variable APP_HTTP_PORT in .env file to convenient value, default is 89,
application will be available by `localhost:APP_HTTP_PORT`

4. Run application

```bash
docker-compose up -d
```

5. Open browser and go to app address.
 Use Load Example Buttons to change html content or create it by yourself.
 Get pdf file with entered content.