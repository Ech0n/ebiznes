import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { useState } from 'react';
import Products from './components/Products';
import Cart from './components/Cart';
import Payments from './components/Payments';
import { Product } from './types';

function App() {
  const [cart, setCart] = useState<{ product: Product; quantity: number }[]>([]);

  const addToCart = (product: Product) => {
    setCart(prevCart => {
      const existing = prevCart.find(item => item.product.id === product.id);
      if (existing) {
        return prevCart.map(item =>
          item.product.id === product.id
            ? { ...item, quantity: item.quantity + 1 }
            : item
        );
      } else {
        return [...prevCart, { product, quantity: 1 }];
      }
    });
  };

  const updateQuantity = (productId: number, newQuantity: number) => {
    setCart(prevCart =>
      prevCart.map(item =>
        item.product.id === productId
          ? { ...item, quantity: newQuantity }
          : item
      )
    );
  };

  const clearCart = () => {
    console.log("clearuing cart")
    setCart([]);
  };

  return (
    <Router>
      <div style={{ padding: '20px' }}>
        <nav style={{
          position: 'fixed',
          top: 0,
          left: 0,
          right: 0,
          padding: '10px 20px',
          marginBottom: '20px',
          background: 'white',
          boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
          zIndex: 1000,
          display: 'flex',
          justifyContent: 'flex-end',
        }}>
          <Link to="/" style={{
            marginRight: '10px',
            color: 'black',
            textDecoration: 'none',
          }}>Sklep</Link>
          <Link to="/cart" style={{
            marginRight: '10px',
            color: 'black',
            textDecoration: 'none',
          }}>Koszyk</Link>
        </nav>

        <Routes>
          <Route path="/" element={<Products onAddToCart={addToCart} />} />
          <Route path="/cart" element={<Cart cart={cart} updateQuantity={updateQuantity} />} />
          <Route path="/payment" element={<Payments cart={cart} clearCart={clearCart} />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
