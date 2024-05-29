import express from 'express';

import { Redis } from './redis';
import { GetConfig } from './config';
import { GetRouter } from './router';
import { ServiceDiscovery } from './service';

const StartWebServer = async () => {
    const config = GetConfig();

    const redis = Redis.Instance({ url: config.redisUrl });
    await redis.connect();

    const service = new ServiceDiscovery(redis.Client());

    const app = express();

    app.use(express.json());
    app.use(express.urlencoded({ extended: true }));

    app.use((req, res, next) => {
        const secretHeader = req.headers['x-secret'];

        if (secretHeader !== config.secret) {
            res.status(401).json({ error: 'Unauthorized' });
            return;
        }

        next();
    });

    const router = GetRouter(service);

    app.use('/api/v1/services', router);
    app.get('/api/v1/health', async (req, res) => {
        const health = await redis.healthCheck();

        let status: 'UP' | 'DOWN' = 'UP';

        if (health.status === 'DOWN') {
            status = 'DOWN';
        }

        res.json({
            status,
            redis: health,
        });
    });

    app.get('*', (req, res) => {
        res.status(404).json({ error: 'Not found' });
    });

    app.listen(config.port, () => {
        console.log('Server is running on port 3000');
    });
};

export { StartWebServer };
