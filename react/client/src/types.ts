export interface Product {
  id: number;
  name: string;
  price: number;
  categoryID: number;
}

export interface OrderItem {
  productId: number;
  quantity: number;
}

export interface OrderResponse {
  totalPrice: number;
  orderItems: {
    productId: number;
    quantity: number;
  }[];
}