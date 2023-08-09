import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import OrderList from './components/OrderList';
import ErrorPage from './components/ErrorPage';
import Genres from './components/Genres';
import CreateDish from './components/CreateDish';
import UpdateDish from './components/UpdateDish';
import Home from './components/Home';
import Login from './components/Login';
import Restaurants from './components/Restaurants';
import Menus from './components/Menus';
import Dishes from './components/Dishes';
import OrderDetails from './components/OrderDetails';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {index: true, element: <Home /> },
      {
        path: "/restaurants",
        element: <Restaurants />,
      },
      {
        path: "/restaurants/:id",
        element: <Menus />,
      },
      {
        path: "/dishes/:id",
        element: <Dishes />,
      },
      {
        path: "/dishes/update/:id",
        element: <UpdateDish />,
      },
      {
        path: "/genres",
        element: <Genres />,
      },
      {
        path: "/admin/orders",
        element: <OrderList />,
      },
      {
        path: "/admin/orders/:id",
        element: <OrderDetails />,
      },
      {
        path: "/admin/orders/create",
        element: <CreateDish />,
      },
      {
        path: "/login",
        element: <Login />,
      },
    ]
  }
])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
