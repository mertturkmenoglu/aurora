import { getRandomBool, getRandomElementFromArray, getRandomFloat, getRandomInt } from './random';
import { brands } from './brands';
import { categories } from './categories';
import { products } from './products';
import { productStylesArr } from './styles';
import { productSizesArr } from './sizes';
import { faker } from '@faker-js/faker';

const imageBase = 'https://aurora-dev-eu-central-product-images.s3.eu-central-1.amazonaws.com';
const imageIds = products.map((it) => it.id);

function createProductDto(i: number) {
  const variants = [];
  const variantCount = getRandomInt(3, 10);

  for (let j = 0; j < variantCount; j++) {
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
      styleId: getRandomElementFromArray(productStylesArr),
      sizeId: getRandomElementFromArray(productSizesArr),
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
    await fetch('http://localhost:5000/api/v1/products', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(dto),
    });
    console.log(`Product ${i + 1} created!`);
  }
}

createProducts();