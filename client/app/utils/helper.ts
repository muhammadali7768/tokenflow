export const weiToEther=(wei:any) =>{
    const ether = wei / Math.pow(10, 18);
    return ether;
  }