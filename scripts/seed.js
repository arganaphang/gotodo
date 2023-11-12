import { check } from "k6";
import http from "k6/http";
import Faker from "https://cdnjs.cloudflare.com/ajax/libs/Faker/3.1.0/faker.min.js";

export const options = {
  scenarios: {
    seed: {
      executor: "constant-arrival-rate",
      duration: "10s",
      preAllocatedVUs: 10,
      rate: 10,
      timeUnit: "1s",
    },
  },
};

export default function () {
  const title = Faker.lorem.words();
  const res = http.post(
    "http://localhost:8000/todos",
    JSON.stringify({ title }),
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  check(res, {
    "Post status is 201": (r) => res.status === 201,
    "Post Content-Type header": (r) =>
      res.headers["Content-Type"] === "application/json; charset=utf-8",
    "Post response name": (r) =>
      res.status === 201 && res.json().data.title === title,
  });
}
