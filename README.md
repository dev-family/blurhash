# Blurhash

BlurHash applies a simple DCT transformation to the image data, 
retaining only the first few components, and then encodes these components using base 83 encoding with the 
JSON character set, HTML:

```
eTGu,l%Mxu-;t7~qxut7t7ofRjWBt7IUWBRjofWBRjj]-;WCIoxut7
```
You can then perform a reverse conversion to sRGB space over the string.

You can read more about the algorithm [here](https://github.com/woltapp/blurhash/blob/master/Algorithm.md).

## Example of image conversion

![img_1.png](img_1.png)            ![img_2.png](img_2.png)

## Service life cycle

You make an http-request to the service, passing in the url of the request a link to your image, as well as x and y parameters (values from 1 to 9) - the number of displayed components.

## How does it work?

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


## How to deploy the service

### docker


```
docker run -d -p 3333:3333 ghcr.io/dev-family/blurhash:1.0.0
```


### docker-compose

```yaml
blurhash:
  image: ghcr.io/dev-family/blurhash:1.0.0
  ports:
    - "3333:3333"
```


