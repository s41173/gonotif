import http from 'k6/http';

export default function () {
  const payload = JSON.stringify({
    username: '082277014410',
    password: 'j4ykiran'
  });

  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const res = http.post('http://localhost:8080/login', payload, params);

  console.log(res.status);
  console.log(res.body);
}