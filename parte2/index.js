let values = [102, 99.2, 109, 113];

const betterDays = (array) => {
  let betterSale = array[0];
  let betterBuy = array[0];

  array.map((Element, index) => {
    if (Element > betterBuy) {
      betterBuy = index;
    }
    if (Element < betterSale) {
      betterSale = index;
    }
  });
  result = [betterBuy, betterSale];
  return result;
};






console.log(betterDays(values));
