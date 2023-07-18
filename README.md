# Blurhash

BlurHash берет ваше изображение и преобразует его в короткую строку (всего 20-30 символов), которая представляет собой размытый фон вашего фото

## Жизненный цикл сервиса

Вы делаете http-запрос на сервис, передавая в урл запроса ссылку на вашу картинку, а также x и y парметры (значения от 1 до 9), от которых зависит качество итогового изображения

## Пример

#### Request

```
http://localhost:3333?image_url=https://imgproxy.com/imageurl&x=1&y=4
```

#### Response

```json
{
  "blur_hash": "eTGu,l%Mxu-;t7~qxut7t7ofRjWBt7IUWBRjofWBRjj]-;WCIoxut7"
}
```


## Как развернуть сервис

### docker


```
docker run -d -p 3333:3333 ghcr.io/dev-family/blurhash
```


### docker-compose

```yaml
blurhash:
  image: ghcr.io/dev-family/blurhash:1.0.0
  ports:
    - "3333:3333"
```


