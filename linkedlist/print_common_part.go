package linkedlist

import "fmt"

func PrintCommonPart(head1, head2 *Node) {
	for head1 != nil && head2 != nil {
		if head1.value < head2.value {
			head1 = head1.next
		} else if head1.value > head2.value {
			head2 = head2.next
		} else {
			fmt.Printf("%d ", head1.value)
			head1 = head1.next
			head2 = head2.next
		}
	}
	fmt.Println()
}
