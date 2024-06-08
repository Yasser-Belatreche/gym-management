import { Router } from 'express';

import { ServiceDiscovery } from './service';

const GetRouter = (service: ServiceDiscovery): Router => {
    const router = Router();

    router.get('/', async (req, res): Promise<void> => {
        const services = await service.getServices();

        res.json(services);
    });

    router.get('/:name', async (req, res): Promise<void> => {
        const { name } = req.params;

        const result = await service.getService(name);

        if (!result) {
            res.status(404).json({ error: 'Service not found' });
            return;
        }

        res.json(result);
    });

    router.post('/', async (req, res): Promise<void> => {
        const { name, url } = req.body;

        if (!name || !url) {
            res.status(400).json({ error: 'Invalid name or URL' });
            return;
        }

        try {
            const s = await service.register(name, url);

            res.json(s);
        } catch (error) {
            res.status(400).json({ error: (error as Error).message });
        }
    });

    router.delete('/:name', async (req, res): Promise<void> => {
        const { name } = req.params;

        await service.deleteService(name);

        res.json({ message: 'Service deleted' });
    });

    return router;
};

export { GetRouter };
