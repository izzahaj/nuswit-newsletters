import { Router } from "express";
import eventController from "../controllers/eventController";

const eventRouter = Router();

eventRouter.get("/events", eventController.getAllEvents);
eventRouter.post("/events", eventController.createEvent);
eventRouter.get("/events/:id", eventController.getEventById);
eventRouter.patch("/events/:id", eventController.updateEvent);
eventRouter.delete("/events/:id", eventController.deleteEvent);

export default eventRouter;
