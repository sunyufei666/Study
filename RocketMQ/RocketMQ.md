# 消息通信机制及实现框架

但凡在业务需求中出现如“当...发生...时、一旦出现...”等描述时，就应该考虑是否需要在这些场景中引入事件。

事件和事件驱动架构

所谓事件，就是将系统中所发生的业务状态变更抽取出来形成一系列独立的对象。而关于如何在系统中生成、发布以及消费事件，业界也存在一个非常重要的设计模式，即事件驱动架构（Event-Driven Architecture, EDA）模式。

四大核心优势：分布式解耦、系统扩展、流量削峰、数据最终一致性

![image-20260920103724415](./images/image-20260920103724415.png)

消息发布：普通消息、顺序消息、延迟消息、事务消息、单向消息、批量消息

消息消费：拉（pull）模式消费、推（push）模式消费、消息过滤（Filter）

![image-20260920104828787](./images/image-20260920104828787.png)

延迟消息、事务消息、消息过滤、消息查询等功能只有RocketMQ支持。

# RocketMQ基本概念和架构

消息是RocketMQ生产和消费数据的最小单位，每条消息必须属于一个主题。

```java
public class Message implements Serializable {
	private String topic;
	private int flag;
	private Map<String, String> properties; // 包含Keys、Tags、UserProperty、DelayTimeLevel等
	private byte[] body;
	private String transactionId;
}
```

主题（Topic）表示一类消息的集合，每个主题包含若干条消息，每条消息只能属于一个主题，是RocketMQ进行消息订阅的基本单位；一个生产者可以同时发送多种Topic的消息；而一个消费者只对某种特定的Topic感兴趣，即只可以订阅和消费一种Topic的消息。

队列（Queue）是存储消息的物理实体，一个Topic中可以包含多个queue，每个queue中存放的就是该Topic的消息，也被称为一个Topic中消息的分区（Partition），一个Topic的queue中的消息只能被一个消费者组中的一个消费者消费，一个queue中的消息不允许同一个消费者组中的多个消费者同时消费。

标签（Tag）是为消息设置的标签，用于同一主题下区分不同类型的消息，来自同一业务单元的消息，可以根据不同业务目的在同一主题下设置不同标签。标签能够有效地保持代码的清晰度和连贯性，并优化RocketMQ提供的查询系统。消费者可以根据Tag实现对不同子主题的不同消费逻辑，实现更好的扩展性。Topic是消息的一级分类，Tag是消息的二级分类。







RocketMQ消息发送方式

RocketMQ消息消费方式