'use client';

import { useQuery } from '@tanstack/react-query';
import { BASE_URL, PaginatedDto, Product } from '@/lib/api';
import ProductCard from '@/components/product-card';

function Page() {
  const { data, isLoading, isError } = useQuery(['products'], async () => {
    const res = await fetch(`${BASE_URL}/products/all?page=1&pageSize=30`);
    return (await res.json()) as PaginatedDto<Product[]>;
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error</div>;
  }

  return (
    <>
      <h1 className="text-4xl font-bold text-center mt-32">Products</h1>
      <div className="container mx-auto grid grid-cols-4 gap-8 mt-32">
        {data &&
          data.data.map((product) => (
            <ProductCard
              product={product}
              key={product.id}
            />
          ))}
      </div>
    </>
  );
}

export default Page;
