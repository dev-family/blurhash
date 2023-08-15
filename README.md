# Blurhash

BlurHash применяет простое преобразование DCT к данным изображения, 
сохраняя только первые несколько компонентов, а затем кодирует эти компоненты, используя кодировку base 83 с 
набором символов JSON, HTML:

```
eTGu,l%Mxu-;t7~qxut7t7ofRjWBt7IUWBRjofWBRjj]-;WCIoxut7
```
Затем над строкой можно выполнить обратное преобразование в пространство sRGB.

Подробнее про алгоритм можно прочитать [тут](https://github.com/woltapp/blurhash/blob/master/Algorithm.md).

## Пример преобразования изображения

![img_1.png](img_1.png)            ![img_2.png](img_2.png)

## Жизненный цикл сервиса

Вы делаете http-запрос на сервис, передавая в урл запроса ссылку на вашу картинку, а также x и y парметры (значения от 1 до 9) - количество отображаемых компонентов.

## Как это работает?

### Request

```
http://localhost:3333?image_url=https://imgproxy.com/imageurl&x=1&y=4
```

### Response

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


