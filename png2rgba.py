from PIL import Image
from tqdm import tqdm

with open('image.bin', 'wb') as f:
    out = bytearray()
    img = Image.open('blog.png')
    rgb_im = img.convert('RGBA')
    #rgb_im = rgb_im.resize((240,160), Image.BICUBIC)
    #rgb_im = rgb_im.crop((0,0,240,160))
    width = rgb_im.size[0]
    height = rgb_im.size[1]
    for pixel in tqdm(range(width*height)):
        #print(pixel)
        x = pixel % width
        y = pixel // width
        r, g, b, a = rgb_im.getpixel((x, y))
        #print(x, y, r, g, b, a)
        pixelBytes = bytearray([x, r, g, b, a])
        for i in pixelBytes:
            out.append(i)
    f.write(bytes(out))