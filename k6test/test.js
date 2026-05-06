import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '30s', target: 20 },  // naik ke 20 user
        { duration: '1m', target: 20 },   // tahan 20 user selama 1 menit
        { duration: '30s', target: 0 },   // turun ke 0
    ],
};

export default function () {
    const res = http.get('http://localhost:8080');
    check(res, {
        'status 200': (r) => r.status === 200,
        'response < 500ms': (r) => r.timings.duration < 500,
    });
    sleep(1);
}