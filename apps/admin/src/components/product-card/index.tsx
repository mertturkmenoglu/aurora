import { useProductMessage } from '@/hooks/useProductMessage';
import { Product } from '@/lib/api';
import { clsx } from 'clsx';
import { useMemo } from 'react';
import Link from 'next/link';

interface Props {
  product: Product;
}

function ProductCard({ product }: Props) {
  const productMessage = useProductMessage(product);
  const image = useMemo(() => {
    if (!product.images.length) {
      return '';
    }

    return product.images[0]?.url ?? '';
  }, [product.images]);

  return (
    <div className={clsx('p-4 flex flex-col rounded-lg group')}>
      <div className="flex items-center justify-between">
        <span className="font-bold text-sm text-sky-600">{productMessage}</span>
      </div>

      <Link
        href={`/products/${product.id}`}
        className={clsx('flex flex-col')}
      >
        <img
          src={image}
          alt=""
          className="w-64 h-48 sm:w-64 md:h-96 object-cover mt-2 mx-auto"
          loading="lazy"
        />
        <div className="mt-2">
          <span className="font-bold text-green-600">{product.currentPrice}$</span>
          {product.currentPrice !== product.oldPrice && (
            <span className="font-light line-through ml-2">{product.oldPrice}</span>
          )}
        </div>
        <span className="text-gray-600 text-sm mt-2">{product.category.name}</span>
        <span className="font-bold mt-2 line-clamp-2 leading-4 h-8">{product.name}</span>
        <span className="font-light text-sm mt-2">By {product.brand.name}</span>
      </Link>
    </div>
  );
}

export default ProductCard;
