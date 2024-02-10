import { Schema, model } from "mongoose";

const EventSchema = new Schema({
  title: {
    type: String,
    required: true
  },
  deadline: Date,
  link: {
    type: String,
    required: true
  },
  venue: String,
  remarks: String
});

const Event = model("Event", EventSchema);

export default Event;
