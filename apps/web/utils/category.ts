import { Category, Product } from '~/utils/dto';

export type ProductPageCategory = Pick<Category, 'id' | 'name'>;

export type ProductPageBreadcrumbLink = {
  name: string;
  href: string;
}

export function getCategoriesFromProduct(product: Product | undefined): ProductPageCategory[] {
  const arr: ProductPageCategory[] = [];

  if (!product) {
    return arr;
  }

  arr.push({
    id: product.category.id,
    name: product.category.name,
  });

  let node = product.category.parent;

  while (node) {
    arr.push({
      id: node.id,
      name: node.name,
    });

    node = node.parent;
  }

  return arr.reverse();
}

export function getBreadcrumbLinksFromCategories(categories: ProductPageCategory[]): ProductPageBreadcrumbLink[] {
  return categories.map((c) => ({
    name: c.name,
    href: `/categories/${c.id}`,
  }));
}
