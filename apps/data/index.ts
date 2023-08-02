import { getRandomBool, getRandomElementFromArray, getRandomFloat, getRandomInt } from './random';
import { brands } from './brands';
import { categories } from './categories';
import { products } from './products';
import { productStylesArr } from './styles';
import { productSizesArr } from './sizes';
import { faker } from '@faker-js/faker';

const ACCESS_TOKEN = '<YOUR_ACCESS_TOKEN>';
const ADMIN_KEY = '<YOUR_ADMIN_KEY>';

const imageBase = 'https://aurora-dev-eu-central-product-images.s3.eu-central-1.amazonaws.com';
const imageIds = products.map((it) => it.id);

function createProductDto(i: number) {
  const variants = [];
  const variantCount = getRandomInt(3, 6);
  console.log(`Creating product ${i + 1} with ${variantCount} variants`);
  const zipped: Array<[string, string]> = [];
  for (let style of productStylesArr) {
    for (let size of productSizesArr) {
      zipped.push([style, size]);
    }
  }

  const chosen = faker.helpers.arrayElements(zipped, variantCount);

  for (let j = 0; j < variantCount; j++) {
    const [style, size] = chosen[j];

    if (!style || !size) {
      throw new Error('Style or size is undefined');
    }

    const currentPrice = getRandomFloat(5, 100);
    const oldPrice = getRandomBool() ? parseFloat((currentPrice + getRandomFloat(1, 10)).toFixed(2)) : currentPrice;
    const v = {
      isDefault: false,
      currentPrice,
      oldPrice,
      inventory: getRandomInt(5, 50),
      shippingPrice: getRandomFloat(0, 20),
      shippingType: 'Direct',
      shippingTime: '2-3 days',
      styleId: style,
      sizeId: size,
      image: {
        url: `${imageBase}/${getRandomElementFromArray(imageIds)}.jpg`,
      },
    };

    variants.push(v);
  }

  variants[0].isDefault = true;

  return {
    name: products[i].name,
    description: faker.commerce.productDescription(),
    isFeatured: getRandomBool(),
    isNew: true,
    isOnSale: true,
    isPopular: getRandomBool(),
    brandId: getRandomElementFromArray(brands),
    categoryId: getRandomElementFromArray(categories),
    variants,
  };
}

async function createProducts() {
  for (let i = 0; i < products.length; i++) {
    const dto = createProductDto(i);
    const res = await fetch('http://localhost:5000/api/v1/products', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'x-access-token': ACCESS_TOKEN,
        'x-admin-key': ADMIN_KEY,
      },
      body: JSON.stringify(dto),
    });

    if (res.ok) {
      console.log(`Product ${i + 1} created!`);
      console.log('-----------------------------------');
      continue;
    }

    console.log(`Product ${i + 1} failed!`);
    console.log('-----------------------------------');
    const body = await res.json();

    console.log(body);
  }
}

createProducts();