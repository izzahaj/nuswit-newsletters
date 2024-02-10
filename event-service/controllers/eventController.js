import eventService from "../services/eventService";

const createEvent = async (req, res) => {
  const { title, deadline, link, venue, remarks } = req.body;

  try {
    const createdEvent = await eventService.createEvent(
      title,
      deadline,
      link,
      venue,
      remarks
    );

    res.status(201).json(createdEvent);
  } catch (err) {
    res.status(500).json({ message: err.message });
  }
}

const getEventById = async (req, res) => {
  const id = req.params.id;

  if (!id) {
    return res.status(400).json({ message: "Missing ID" });
  }

  try {
    const event = await eventService.getEventById(id);

    if (!event) {
      return res.status(404).json({ message: "No such event found" });
    }

    res.status(200).json(event);
  } catch (err) {
    res.status(500).json({ message: err.message });
  }
}

const updateEvent = async (req, res) => {
  const id = req.params.id;
  const { title, deadline, link, venue, remarks } = req.body;

  if (!id) {
    return res.status(400).json({ message: "Missing ID" });
  }

  try {
    const updatedEvent = await eventService.updateEvent(id, req.body);

    if (!updatedEvent) {
      return res.status(404).json({ message: "No such event found" });
    }

    res.status(200).json(updatedEvent);
  } catch (err) {
    res.status(500).json({ message: err.message });
  }
}

const deleteEvent = async (req, res) => {
  const id = req.params.id;

  if (!id) {
    return res.status(400).json({ message: "Missing ID" });
  }

  try {
    await eventService.deleteEvent(id);

    res.status(204).json({ message: `Event with id ${id} has been deleted`});
  } catch (err) {
    res.status(500).json({ message: err.message });
  }
}

const getAllEvents = async (req, res) => {
  try {
    const events = await eventService.getAllEvents();
    res.status(200).json(events);
  } catch (err) {
    res.status(500).json({ message: err.message });
  }
}

export default eventController = { createEvent, getEventById, updateEvent, deleteEvent, getAllEvents };
