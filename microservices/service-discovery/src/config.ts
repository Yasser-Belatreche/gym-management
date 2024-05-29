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

    const port = parseInt(process.env.PORT || '3000');

    if (!process.env.SECRET) throw new Error(`SECRET env variable is required`);

    const secret = process.env.SECRET;

    return { redisUrl, port, secret };
};

export { GetConfig };
