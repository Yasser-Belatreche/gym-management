import { Router } from 'express';

import { ServiceDiscovery } from './service.js';

const GetRouter = (service: ServiceDiscovery): Router => {
    const router = Router();

    router.get('/', async (req, res): Promise<void> => {
        const services = await service.getServices();

        res.json(services);
    });

    router.get('/:id', async (req, res): Promise<void> => {
        const { id } = req.params;

        const result = await service.getService(id);

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

            res.status(201).json(s);
        } catch (error) {
            res.status(400).json({ error: (error as Error).message });
        }
    });

    router.put('/:id', async (req, res): Promise<void> => {
        const { id } = req.params;
        const { name, url } = req.body;

        if (!name || !url) {
            res.status(400).json({ error: 'Invalid name or URL' });
            return;
        }

        try {
            const s = await service.updateService(id, name, url);

            res.json(s);
        } catch (error) {
            const message = (error as Error).message;

            if (message === 'Service not found') {
                res.status(404).json({ error: message });
            } else {
                res.status(400).json({ error: message });
            }
        }
    });

    router.delete('/:id', async (req, res): Promise<void> => {
        const { id } = req.params;

        try {
            await service.deleteService(id);

            res.status(204).end();
        } catch (error) {
            const message = (error as Error).message;

            if (message === 'Service not found') {
                res.status(404).json({ error: message });
            } else {
                res.status(400).json({ error: message });
            }
        }
    });

    return router;
};

export { GetRouter };
