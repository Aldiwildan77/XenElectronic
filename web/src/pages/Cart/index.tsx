import React, { useState, useEffect } from 'react';
import client from '../../lib/client';

import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import LocalMallIcon from '@material-ui/icons/LocalMall';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';
import { makeStyles, CircularProgress } from '@material-ui/core';

import CartItem from '../../lib/model/CartItem';
import Product from '../../lib/model/Product';
import { useCookies } from 'react-cookie';

const useStyles = makeStyles((theme) => ({
  table: {
    midWidth: 700,
  },
  checkoutWrapper: {
    marginBottom: 24,
  },
}))

const currencyFormatter = new Intl.NumberFormat('id-ID', {
  style: 'currency',
  currency: 'IDR',
})

export default function Cart() {
  const classes = useStyles();
  const [cartItems, setCartItems] = useState<CartItem[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isFetched, setIsFetched] = useState<boolean>(false);
  const [error, setError] = useState<Error | null>(null);
  const [{ cartId }] = useCookies(['cartId']);

  useEffect(() => {
    if (!isLoading && !isFetched) {
      setIsLoading(true);
      setError(null);
      client.listProductsInCart(cartId)
        .then(response => response.json())
        .then(response => {
          if (Array.isArray(response)) {
            setCartItems(response.map(cartItem => new CartItem(
              cartItem.id,
              cartItem.productId,
              new Product(
                cartItem.product.id,
                cartItem.product.categoryId,
                cartItem.product.name,
                cartItem.product.description,
                cartItem.product.photo,
                cartItem.product.price,
              )
            )));
          }
        })
        .catch(error => {
          setCartItems([]);
          setError(error)
        })
        .finally(() => {
          setIsLoading(false);
          setIsFetched(true);
        })
    }
  }, [isLoading, isFetched, cartId]);

  const loadingContent = (
    <Grid item xs={12}>
      <CircularProgress />
    </Grid>
  );

  const totalPrice: number = cartItems.reduce((sum, cartItem) => sum + cartItem.product.price, 0)

  const content = (
    <React.Fragment>
      <Grid item xs={12}>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="spanning table">
            <TableHead>
              <TableRow>
                <TableCell>No.</TableCell>
                <TableCell>Name</TableCell>
                <TableCell>Price</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {cartItems.map((cartItem, index) => (
                <TableRow key={cartItem.product.name}>
                  <TableCell>{index + 1}</TableCell>
                  <TableCell align="left">{cartItem.product.name}</TableCell>
                  <TableCell align="left">{currencyFormatter.format(cartItem.product.price)}</TableCell>
                </TableRow>
              ))}
              <TableRow>
                <TableCell />
                <TableCell colSpan={1} align="right">Total</TableCell>
                <TableCell align="left">{currencyFormatter.format(totalPrice)}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      </Grid>
      <Grid item container xs={12} justify="flex-end" className={classes.checkoutWrapper}>
        {/* <Grid container direction="" */}
        <Button
          variant="contained"
          color="primary"
          size="large"
          startIcon={<LocalMallIcon />}
        >
          Checkout
        </Button>
      </Grid>
    </React.Fragment>
  );

  const errorContent = (
    <Grid item xs={12}>
      <Typography variant="body1" color="error">Sorry, something is broken</Typography>
    </Grid>
  );

  return (
    <Grid container spacing={3}>
      {
        isLoading ? loadingContent : (
          error !== null ? errorContent : (
            content
          )
        )
      }
    </Grid>
  );
}