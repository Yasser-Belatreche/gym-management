interface Config {
    redisUrl: string;
    port: number;
    secret: string;
}

const GetConfig = (): Config => {
    const redisUrl = process.env.REDIS_URL;
    if (!redisUrl) {
        throw new Error('REDIS_URL env variable is required');
    }

    const port = parseInt(process.env.PORT ?? '3000');

    if (!process.env.API_SECRET) throw new Error(`API_SECRET env variable is required`);

    const secret = process.env.API_SECRET;

    return { redisUrl, port, secret };
};

export { GetConfig };
