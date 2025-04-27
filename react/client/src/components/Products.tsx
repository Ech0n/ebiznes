import { useEffect, useState } from 'react';
import { api } from '../api';
import { Product } from '../types';

const Products = ({ onAddToCart }: { onAddToCart: (product: Product) => void }) => {
    const [products, setProducts] = useState<Product[]>([]);

    useEffect(() => {
        api.get<Product[]>('/products')
            .then((response) => {
                const mappedProducts: Product[] = response.data.map((item: any) => ({
                    id: item.ID,
                    name: item.Name,
                    price: item.Price,
                    categoryID: item.CategoryID,
                    category: {
                        id: item.Category.ID,
                        name: item.Category.Name,
                    }
                }));
                setProducts(mappedProducts);
                console.log(response.data)
                console.log("products ",products)
            })
            .catch((error)=> console.error('Error fetching products:', error));
    }, []);

    return (
        <div>
            <h2>Produkty</h2>
            <ul>
                {products.map(product => (
                    <li key={product.id}>
                        {product.name} - {product.price} PLN
                        <button onClick={() => onAddToCart(product)} style={{ marginLeft: 10 }}>Dodaj do koszyka</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Products;