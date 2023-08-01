export const BASE_URL = 'http://localhost:5000/api/v1';

export type Dto<T> = {
  data: T;
};

export type Pagination = {
  page: number;
  pageSize: number;
  totalRecords: number;
  totalPages: number;
  hasPrevious: boolean;
  hasNext: boolean;
};

export type PaginatedDto<T> = {
  data: T;
  pagination: Pagination;
};

export type LoginResponseDto = Dto<{
  accessToken: string;
  refreshToken: string;
}>;

export type UserDto = Dto<User>;

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
};

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
};

export type ProductsDto = Dto<Product[]>;
export type ProductDto = Dto<Product>;

export type Product = {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  description: string;
  currentPrice: number;
  oldPrice: number;
  inventory: number;
  images: ProductImage[];
  isFeatured: boolean;
  isNew: boolean;
  isOnSale: boolean;
  isPopular: boolean;
  shippingPrice: number;
  shippingTime: string;
  shippingType: string;
  slug: string;
  brandId: string;
  brand: Brand;
  categoryId: string;
  category: Category;
  styles: ProductStyle[];
  sizes: ProductSize[];
};

export type ProductImage = {
  id: string;
  productId: string;
  url: string;
};

export type ProductStyle = {
  id: string;
  productId: string;
  name: string;
};

export type ProductSize = {
  id: string;
  productId: string;
  name: string;
};

export type Brand = {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  description: string;
};

export type Category = {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  parentId: string | null;
  parent: Category | null;
};

export type Cart = {
  id: string;
  createdAt: string;
  updatedAt: string;
  userId: string;
  items: CartItem[];
};

export type CartItem = {
  id: string;
  createdAt: string;
  updatedAt: string;
  cartId: string;
  productId: string;
  product: Product;
  quantity: number;
};

export type CartDto = Dto<Cart>;
