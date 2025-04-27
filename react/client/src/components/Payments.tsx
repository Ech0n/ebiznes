import { useState } from 'react';
import { api } from '../api';
import { Product,OrderItem } from '../types';

interface PaymentsProps {
    cart: { product: Product; quantity: number }[];
    clearCart: () => void;
}

const Payments = ({ cart, clearCart }: PaymentsProps) => {
    console.log("cart ",cart)
    const total = cart.reduce((sum, item) => sum + item.product.price * item.quantity, 0);
    const [orderConfirmed, setOrderConfirmed] = useState<boolean>(false);
    const handleConfirmOrder = async () => {
        const items: OrderItem[] = cart.map(({ product, quantity }) => ({
            productId: product.id,
            quantity,
        }));

        try {
            await api.post('/payments', { items });
            setOrderConfirmed(true);
            clearCart();
        } catch (error) {
            console.error('Error confirming order:', error);
        }
    };

    return (
        <div>
            <h2>Płatności</h2>
            {cart.length === 0 ? (
                <p>Twój koszyk jest pusty.</p>
            ) : (
                <div>
                    {cart.map(({ product, quantity }) => (
                        <li key={product.id}>
                            {product.name} x {quantity} — {product.price * quantity} PLN
                        </li>
                    ))}
                    <div>
                            <p>Łączna cena: {total} PLN</p>
                        <button onClick={handleConfirmOrder}>Potwierdź i zapłać</button>
                    </div>
                </div>
            )}

            {orderConfirmed && <p>Zamówienie zostało złożone!</p>}
        </div>
    );
};

export default Payments;
