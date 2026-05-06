import http from 'k6/http';
import { check } from 'k6';

export const options = {
  vus: 50,
  duration: '30s',
};

export default function () {
  const payload = JSON.stringify({
    limit: "100",
    offset: "0",
    orderby: "name",
    order: "desc",
    category: "",
    location: "",
    condition: ""
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  // const res = http.post('http://localhost:8080/product', payload, params);
  const res = http.post('https://goapi.dswip.com/product', payload, params);

  check(res, {
    'status is 200': (r) => r.status === 200,
  });
}