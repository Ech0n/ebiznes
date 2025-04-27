import { Product } from '../types';
import { useNavigate } from "react-router-dom";

interface CartProps {
    cart: { product: Product; quantity: number }[];
    updateQuantity: (productId: number, newQuantity: number) => void;
}

const Cart = ({ cart, updateQuantity }: CartProps) => {
    const total = cart.reduce((sum, item) => sum + item.product.price * item.quantity, 0);
    const navigate = useNavigate();

    function handleClick() {
        navigate("/payment");
    }
    return (
        <div>
            <h2>Koszyk</h2>
            {cart.length === 0 ? (
                <p>Koszyk jest pusty.</p>
            ) : (
                <ul>
                    {cart.map(({ product, quantity }) => (
                        <li key={product.id}>
                            {product.name} x {quantity} — {product.price * quantity} PLN
                            <button onClick={() => updateQuantity(product.id, quantity + 1)}>+</button>
                            <button onClick={() => updateQuantity(product.id, quantity - 1)} disabled={quantity <= 1}>-</button>
                        </li>
                    ))}
                </ul>
            )}
            <h3>Łącznie: {total} PLN</h3>
            <button onClick={() => handleClick()}>Przejdź do płatności</button>
        </div>
    );
};

export default Cart;
