package org.projects.cs.frontend.chat.converter;

import org.projects.cs.frontend.chat.entity.ChatRecord;
import org.projects.cs.frontend.chat.event.TicketGeneratedEvent;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;


@Mapper
public interface ChatRecordConverter {
    ChatRecordConverter INSTANCE = Mappers.getMapper(ChatRecordConverter.class);

    //Event->Entity
    ChatRecord convertEvent(TicketGeneratedEvent event);
}
