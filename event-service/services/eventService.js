import Event from "../models/eventModel";

const createEvent = async (
  title, deadline, link, venue, remarks
) => {
  const newEvent = new Event({
    title,
    deadline,
    link,
    venue,
    remarks
  });

  return await newEvent.save();
}

const getEventById = async (id) => {
  return await Event.findById(id);
}

const updateEvent = async (id, updatedEventData) => {
  const options = { new: true };
  return await Event.findByIdAndUpdate(id, updatedEventData, options);
}

const deleteEvent = async (id) => {
  return await Event.findByIdAndDelete(id);
}

const getAllEvents = async () => {
  return await Event.find().sort({ deadline: 1 });
}

export default eventService = { createEvent, getEventById, updateEvent, deleteEvent, getAllEvents };
