import express from 'express'
import morgan from 'morgan'
import postgraphql from 'postgraphql'

const app = express()
app
    .use(morgan(`combined`))
    .use(postgraphql.postgraphql(
        {user: `jordan`, host: `localhost`, database: `postgres`},
        `dashboard`,
        {dynamicJson: true, graphiql: true, watchPg: true},
    ))
    .listen(3000)

console.log(`app is listening on port 3000`);
