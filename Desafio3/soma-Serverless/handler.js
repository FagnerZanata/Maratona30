"use strict";

module.exports.soma = async event => {
  const query = event.queryStringParameters;
  let body;

  if (!query)
    body = {
      error: { message: "Parâmetros invalidos" },
    };
  else {
    const a = Number(query.a);
    const b = Number(query.b);

    body =
      !a || !b
        ? { error: { message: "Parâmetros com valores invalidos" } }
        : { resultado: a + b };
  }

  return {
    statusCode: 200,
    body: JSON.stringify(body, null, 2),
  };
};