import http from 'k6/http';
import { check, fail, sleep } from 'k6';

import environment from 'k6/x/setenv';
import lgcPinger from 'k6/x/pinger';

export const options = {
    stages: [
        { duration: '5s', target: 20 },
        { duration: '10', target: 10 },
        { duration: '5', target: 0 },
    ],
};

export function setup() {
    environment.loadEnvironment();
    lgcPinger.setNewPinger(environment.getEnvironmentVariable("BASE_URL"));
    lgcPinger.runPinger(2);
}

export default function () {
    const res =
        http.get(`https://${environment.getEnvironmentVariable("BASE_URL")}/todos/1`);
    check(res, { 'status was 200': (r) => r.status == 200 });
    sleep(1);
}
