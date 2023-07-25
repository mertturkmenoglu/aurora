import {Product} from "~/utils/dto";

export function useProductMessage(product: Product) {
    if (product.shippingPrice === 0) {
        return 'Free Shipping'
    } else if (product.isPopular) {
        return 'Popular'
    } else if (product.isNew) {
        return 'New'
    } else if (product.isOnSale) {
        return 'On Sale'
    } else {
        return ''
    }
}