export type Dto<T> = {
  data: T
}

export type LoginResponseDto = Dto<{
  accessToken: string
  refreshToken: string
}>

export type UserDto = Dto<User>

export type User = {
  id: string;
  fullName: string;
  email: string;
  phone: string;
  adPreference: AdPreference;
  addresses: Address[];
};

export type AdPreference = {
  id: string;
  email: boolean;
  sms: boolean;
  phone: boolean;
}

export type Address = {
  id: string;
  city: string;
  description: string;
  isDefault: boolean;
  line1: string;
  line2: string;
  name: string;
  phone: string;
  state: string;
  type: string;
  zipCode: string;
}

export type ProductsDto = Dto<Product[]>
export type ProductDto = Dto<Product>

export type Product = {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  description: string;
  isFeatured: boolean;
  isNew: boolean;
  isOnSale: boolean;
  isPopular: boolean;
  brandId: string;
  brand: Brand;
  categoryId: string;
  category: Category;
  defaultVariantId: string;
  defaultVariant: ProductVariant;
  variants: ProductVariant[];
}

export type ProductVariant = {
  id: string;
  createdAt: string;
  updatedAt: string;
  productId: string;
  currentPrice: number;
  oldPrice: number;
  inventory: number;
  imageId: string;
  image: ProductImage;
  shippingPrice: number;
  shippingTime: string;
  shippingType: string;
  styleId: string;
  style: ProductStyle;
  sizeId: string;
  size: ProductSize;
}

export type ProductImage = {
  id: string;
  productId: string;
  productVariantId: string;
  url: string;
}

export type ProductStyle = {
  id: string;
  name: string;
}

export type ProductSize = {
  id: string;
  name: string;
}

export type Brand = {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  description: string;
}

export type Category = {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  parentId: string | null;
  parent: Category | null;
}

export type Cart = {
  id: string;
  createdAt: string;
  updatedAt: string;
  userId: string;
  items: CartItem[];
}

export type CartItem = {
  id: string;
  createdAt: string;
  updatedAt: string;
  cartId: string;
  productId: string;
  product: Product;
  quantity: number;
}

export type CartDto = Dto<Cart>

export type CfBanner = {
  title: string;
  href: string;
  enabledPages: string[];
  image: {
    fields: {
      title: string;
      description: string;
      file: {
        url: string;
        details: {
          size: number;
          image: {
            width: number;
            height: number;
          };
        };
        fileName: string;
        contentType: string;
      };
    };
  };
}

export type HomeAggregation = {
  featured: Product[];
  new: Product[];
  sale: Product[];
  popular: Product[];
};

export type HomeAggregationDto = Dto<HomeAggregation>;

