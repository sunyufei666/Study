package org.projects.cs.frontend.ticket.converter;

import org.projects.cs.frontend.ticket.entity.CustomerTicket;
import org.projects.cs.frontend.ticket.event.TicketGeneratedEvent;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;


@Mapper
public interface CustomerTicketConverter {
    CustomerTicketConverter INSTANCE = Mappers.getMapper(CustomerTicketConverter.class);

    //Event->Entity
    CustomerTicket convertEvent(TicketGeneratedEvent event);
}
